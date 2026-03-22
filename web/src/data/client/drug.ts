import {
  PatientQueryOptions,
  QueryOptions,
  Drug,
  CreateDrugInput,
  DrugPaginator,
} from '@/types';
import { API_ENDPOINTS } from './api-endpoints';
import { crudFactory } from './curd-factory';
import { HttpClient } from './http-client';

export const drugClient = {
  ...crudFactory<Drug, QueryOptions, CreateDrugInput>(
    API_ENDPOINTS.DRUGS,
  ),
  paginated: ({ name, ...params }: Partial<PatientQueryOptions>) => {
    return HttpClient.get<DrugPaginator>(API_ENDPOINTS.DRUGS, {
      searchJoin: 'and',
      self,
      ...params,
      search: HttpClient.formatSearchParams({ name }),
    });
  },
};
