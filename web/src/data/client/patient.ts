import {
  Patient,
  PatientPaginator,
  PatientQueryOptions,
  CreatePatientInput,
  QueryOptions,
} from '@/types';
import { API_ENDPOINTS } from './api-endpoints';
import { crudFactory } from './curd-factory';
import { HttpClient } from './http-client';

export const patientClient = {
  ...crudFactory<Patient, QueryOptions, CreatePatientInput>(
    API_ENDPOINTS.PATIENTS,
  ),
  paginated: ({ name, ...params }: Partial<PatientQueryOptions>) => {
    return HttpClient.get<PatientPaginator>(API_ENDPOINTS.PATIENTS, {
      searchJoin: 'and',
      self,
      ...params,
      search: HttpClient.formatSearchParams({ name }),
    });
  },
};
