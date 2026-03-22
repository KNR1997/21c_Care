import Router, { useRouter } from 'next/router';
import { useMutation, useQuery, useQueryClient } from 'react-query';
import { toast } from 'react-toastify';
import { useTranslation } from 'next-i18next';
import { Routes } from '@/config/routes';
import { API_ENDPOINTS } from './client/api-endpoints';
import { Visit, VisitPaginator, VisitQueryOptions, GetParams } from '@/types';
import { mapPaginatorData } from '@/utils/data-mappers';
import { visitClient } from './client/visit';
import { Config } from '@/config';

export const useVisitsQuery = (options: Partial<VisitQueryOptions>) => {
  const { data, error, isLoading } = useQuery<VisitPaginator, Error>(
    [API_ENDPOINTS.VISITS, options],
    ({ queryKey, pageParam }) =>
      visitClient.paginated(Object.assign({}, queryKey[1], pageParam)),
    {
      keepPreviousData: true,
    },
  );

  return {
    visits: data?.data ?? [],
    paginatorInfo: mapPaginatorData(data),
    error,
    loading: isLoading,
  };
};

export const useVisitQuery = ({ slug, language }: GetParams) => {
  const { data, error, isLoading } = useQuery<Visit, Error>(
    [API_ENDPOINTS.CATEGORIES, { slug, language }],
    () => visitClient.get({ slug, language }),
  );

  return {
    visit: data,
    error,
    isLoading,
  };
};

export const useCreateVisitMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(visitClient.create, {
    onSuccess: () => {
      Router.push(Routes.visit.list, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-created'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.VISITS);
    },
  });
};

export const useUpdateVisitMutation = () => {
  const { t } = useTranslation();
  const router = useRouter();
  const queryClient = useQueryClient();
  return useMutation(visitClient.update, {
    onSuccess: async (data) => {
      const generateRedirectUrl = router.query.shop
        ? `/${router.query.shop}${Routes.visit.list}`
        : Routes.visit.list;
      await router.push(`${generateRedirectUrl}/${data?.id}/edit`, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-updated'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.CATEGORIES);
    },
  });
};

export const useDeleteVisitMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(visitClient.delete, {
    onSuccess: () => {
      toast.success(t('common:successfully-deleted'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.VISITS);
    },
  });
};

export const usePreviewVisitMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(visitClient.preview, {
    onSuccess: () => {
      // Router.push(Routes.visit.list, undefined, {
      //   locale: Config.defaultLanguage,
      // });
      toast.success(t('common:successfully-created'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      // queryClient.invalidateQueries(API_ENDPOINTS.VISITS);
    },
  });
};

export const useDownloadInvoiceMutation = (visit_id: number) => {
  return useQuery<Blob, Error>(
    [API_ENDPOINTS.ORDER_INVOICE_DOWNLOAD, visit_id],
    () => visitClient.downloadInvoice(visit_id),
    {
      enabled: false,
    },
  );
};
