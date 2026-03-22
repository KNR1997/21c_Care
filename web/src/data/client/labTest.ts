import {
  PatientQueryOptions,
  QueryOptions,
  CreateLabTestInput,
  LabTest,
  LabTestPaginator,
} from '@/types';
import { API_ENDPOINTS } from './api-endpoints';
import { crudFactory } from './curd-factory';
import { HttpClient } from './http-client';

export const labTestClient = {
  ...crudFactory<LabTest, QueryOptions, CreateLabTestInput>(
    API_ENDPOINTS.LABTESTS,
  ),
  paginated: ({ name, ...params }: Partial<PatientQueryOptions>) => {
    return HttpClient.get<LabTestPaginator>(
      API_ENDPOINTS.LABTESTS  ,
      {
        searchJoin: 'and',
        self,
        ...params,
        search: HttpClient.formatSearchParams({ name }),
      },
    );
  },
};
