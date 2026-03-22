import Router, { useRouter } from 'next/router';
import { useMutation, useQuery, useQueryClient } from 'react-query';
import { toast } from 'react-toastify';
import { useTranslation } from 'next-i18next';
import { Routes } from '@/config/routes';
import { API_ENDPOINTS } from './client/api-endpoints';
import {
  LabTest,
  LabTestPaginator,
  LabTestQueryOptions,
  GetParams,
} from '@/types';
import { mapPaginatorData } from '@/utils/data-mappers';
import { Config } from '@/config';
import { labTestClient } from './client/labTest';

export const useLabTestsQuery = (options: Partial<LabTestQueryOptions>) => {
  const { data, error, isLoading } = useQuery<LabTestPaginator, Error>(
    [API_ENDPOINTS.LABTESTS, options],
    ({ queryKey, pageParam }) =>
      labTestClient.paginated(Object.assign({}, queryKey[1], pageParam)),
    {
      keepPreviousData: true,
    },
  );

  return {
    labTests: data?.data ?? [],
    paginatorInfo: mapPaginatorData(data),
    error,
    loading: isLoading,
  };
};

export const useLabTestQuery = ({ slug, language }: GetParams) => {
  const { data, error, isLoading } = useQuery<LabTest, Error>(
    [API_ENDPOINTS.LABTESTS, { slug, language }],
    () => labTestClient.get({ slug, language }),
  );

  return {
    labTest: data,
    error,
    isLoading,
  };
};

export const useCreateLabTestMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(labTestClient.create, {
    onSuccess: () => {
      Router.push(Routes.labTest.list, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-created'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.LABTESTS);
    },
  });
};

export const useUpdateLabTestMutation = () => {
  const { t } = useTranslation();
  const router = useRouter();
  const queryClient = useQueryClient();
  return useMutation(labTestClient.update, {
    onSuccess: async (data) => {
      const generateRedirectUrl = router.query.shop
        ? `/${router.query.shop}${Routes.labTest.list}`
        : Routes.labTest.list;
      await router.push(`${generateRedirectUrl}/${data?.id}/edit`, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-updated'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.LABTESTS);
    },
  });
};

export const useDeleteLabTestMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(labTestClient.delete, {
    onSuccess: () => {
      toast.success(t('common:successfully-deleted'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.LABTESTS);
    },
  });
};
