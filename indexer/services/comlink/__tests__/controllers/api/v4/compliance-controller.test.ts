import { RequestMethod } from '../../../../src/types';
import { sendRequest } from '../../../helpers/helpers';
import {
  ComplianceDataFromDatabase,
  ComplianceTable,
  dbHelpers,
  testConstants,
  testMocks,
} from '@dydxprotocol-indexer/postgres';
import { stats } from '@dydxprotocol-indexer/base';
import { complianceProvider } from '../../../../src/helpers/compliance/compliance-clients';
import { ComplianceClientResponse } from '@dydxprotocol-indexer/compliance';
import { ratelimitRedis } from '../../../../src/caches/rate-limiters';
import { redis } from '@dydxprotocol-indexer/redis';
import { DateTime } from 'luxon';
import config from '../../../../src/config';
import { getIpAddr } from '../../../../src/lib/rate-limit';

jest.mock('../../../../src/lib/rate-limit', () => ({
  ...jest.requireActual('../../../../src/lib/rate-limit'),
  getIpAddr: jest.fn(),
}));

describe('compliance-controller#V4', () => {
  const riskScore: string = '10.00';
  const blocked: boolean = false;
  const ipAddr: string = '192.168.1.1';

  const ipAddrMock: jest.Mock = (getIpAddr as unknown as jest.Mock);

  beforeAll(async () => {
    await dbHelpers.migrate();
    jest.spyOn(stats, 'increment');
    jest.spyOn(stats, 'timing');
  });

  afterAll(async () => {
    await dbHelpers.teardown();
  });

  describe('GET', () => {
    beforeEach(async () => {
      jest.spyOn(complianceProvider.client, 'getComplianceResponse').mockImplementation(
        (address: string): Promise<ComplianceClientResponse> => {
          return Promise.resolve({
            address,
            blocked,
            riskScore,
          });
        },
      );
      ipAddrMock.mockReturnValue(ipAddr);
      await testMocks.seedData();
    });

    afterEach(async () => {
      await redis.deleteAllAsync(ratelimitRedis.client);
      await dbHelpers.clearData();
    });

    it('Get /screen with new address gets and stores compliance data from provider', async () => {
      let data: ComplianceDataFromDatabase[] = await ComplianceTable.findAll({}, [], {});
      expect(data).toHaveLength(0);

      const response: any = await sendRequest({
        type: RequestMethod.GET,
        path: `/v4/screen?address=${testConstants.defaultAddress}`,
      });

      expect(response.body.restricted).toEqual(false);
      expect(stats.timing).toHaveBeenCalledTimes(1);
      expect(complianceProvider.client.getComplianceResponse).toHaveBeenCalledTimes(1);

      data = await ComplianceTable.findAll({}, [], {});
      expect(data).toHaveLength(1);
      expect(data[0]).toEqual(expect.objectContaining({
        address: testConstants.defaultAddress,
        provider: complianceProvider.provider,
        blocked,
        riskScore,
      }));
    });

    it('Get /screen with existing address retrieves compliance data from database', async () => {
      // Seed the database with a compliance record
      await ComplianceTable.create(testConstants.nonBlockedComplianceData);
      let data: ComplianceDataFromDatabase[] = await ComplianceTable.findAll({}, [], {});
      expect(data).toHaveLength(1);

      const response: any = await sendRequest({
        type: RequestMethod.GET,
        path: `/v4/screen?address=${testConstants.defaultAddress}`,
      });

      expect(response.body.restricted).toEqual(false);
      expect(stats.timing).toHaveBeenCalledTimes(1);
      expect(complianceProvider.client.getComplianceResponse).toHaveBeenCalledTimes(0);

      data = await ComplianceTable.findAll({}, [], {});
      expect(data).toHaveLength(1);
      expect(data[0]).toEqual(testConstants.nonBlockedComplianceData);
    });

    it('Get /screen with old existing address refreshes compliance data', async () => {
      // Seed the database with an old compliance record
      await ComplianceTable.create({
        ...testConstants.nonBlockedComplianceData,
        updatedAt: DateTime.utc().minus({
          seconds: config.MAX_AGE_SCREENED_ADDRESS_COMPLIANCE_DATA_SECONDS * 2,
        }).toISO(),
      });
      let data: ComplianceDataFromDatabase[] = await ComplianceTable.findAll({}, [], {});
      expect(data).toHaveLength(1);

      const response: any = await sendRequest({
        type: RequestMethod.GET,
        path: `/v4/screen?address=${testConstants.defaultAddress}`,
      });

      expect(response.body.restricted).toEqual(false);
      expect(stats.timing).toHaveBeenCalledTimes(1);
      expect(complianceProvider.client.getComplianceResponse).toHaveBeenCalledTimes(1);

      data = await ComplianceTable.findAll({}, [], {});
      expect(data).toHaveLength(1);
      expect(data[0]).not.toEqual({
        address: testConstants.defaultAddress,
        provider: complianceProvider.provider,
        blocked,
        riskScore,
      });
    });

    it('Get /screen with multiple new address from same IP gets rate-limited', async () => {
      for (let i: number = 0; i < config.RATE_LIMIT_SCREEN_QUERY_PROVIDER_POINTS; i++) {
        await sendRequest({
          type: RequestMethod.GET,
          path: `/v4/screen?address=${i}`,
        });
      }

      await sendRequest({
        type: RequestMethod.GET,
        path: `/v4/screen?address=${testConstants.defaultAddress}`,
        errorMsg: 'Too many requests',
        expectedStatus: 429,
      });
    });

    it('Get /screen with multiple new address globally gets rate-limited', async () => {
      ipAddrMock.mockImplementation(() => Math.random().toString());
      for (let i: number = 0; i < config.RATE_LIMIT_SCREEN_QUERY_PROVIDER_GLOBAL_POINTS; i++) {
        await sendRequest({
          type: RequestMethod.GET,
          path: `/v4/screen?address=${i}`,
        });
      }

      await sendRequest({
        type: RequestMethod.GET,
        path: `/v4/screen?address=${testConstants.defaultAddress}`,
        errorMsg: 'Too many requests',
        expectedStatus: 429,
      });
    });
  });
});
