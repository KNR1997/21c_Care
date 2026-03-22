import {
  Billing,
  BillingPaginator,
  BillingQueryOptions,
  CreateBillingInput,
  QueryOptions,
} from '@/types';
import { API_ENDPOINTS } from './api-endpoints';
import { crudFactory } from './curd-factory';
import { HttpClient } from './http-client';

export const billingClient = {
  ...crudFactory<Billing, QueryOptions, CreateBillingInput>(
    API_ENDPOINTS.BILLINGS,
  ),
  paginated: ({ name, ...params }: Partial<BillingQueryOptions>) => {
    return HttpClient.get<BillingPaginator>(API_ENDPOINTS.BILLINGS, {
      searchJoin: 'and',
      self,
      ...params,
      search: HttpClient.formatSearchParams({ name }),
    });
  },
};
