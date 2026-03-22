import {
  Visit,
  VisitPaginator,
  VisitQueryOptions,
  CreateVisitInput,
  QueryOptions,
  PreviewVisitInput,
  PreviewVisitResponse,
  GenerateInvoiceDownloadUrlInput,
} from '@/types';
import { API_ENDPOINTS } from './api-endpoints';
import { crudFactory } from './curd-factory';
import { HttpClient } from './http-client';

export const visitClient = {
  ...crudFactory<Visit, QueryOptions, CreateVisitInput>(API_ENDPOINTS.VISITS),
  paginated: ({ name, ...params }: Partial<VisitQueryOptions>) => {
    return HttpClient.get<VisitPaginator>(API_ENDPOINTS.VISITS, {
      searchJoin: 'and',
      self,
      ...params,
      search: HttpClient.formatSearchParams({ name }),
    });
  },
  preview: (data: PreviewVisitInput) => {
    return HttpClient.post<PreviewVisitResponse>(
      `${API_ENDPOINTS.VISITS}/preview`,
      data,
    );
  },
  downloadInvoice: (visit_id: number) => {
    return HttpClient.get<Blob>(`${API_ENDPOINTS.VISITS}/${visit_id}/report`, {
      responseType: 'blob',
    });
  },
};
