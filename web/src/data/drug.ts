import Router, { useRouter } from 'next/router';
import { useMutation, useQuery, useQueryClient } from 'react-query';
import { toast } from 'react-toastify';
import { useTranslation } from 'next-i18next';
import { Routes } from '@/config/routes';
import { API_ENDPOINTS } from './client/api-endpoints';
import {
  Drug,
  DrugPaginator,
  DrugQueryOptions,
  GetParams,
} from '@/types';
import { mapPaginatorData } from '@/utils/data-mappers';
import { Config } from '@/config';
import { drugClient } from './client/drug';

export const useDrugsQuery = (options: Partial<DrugQueryOptions>) => {
  const { data, error, isLoading } = useQuery<DrugPaginator, Error>(
    [API_ENDPOINTS.DRUGS, options],
    ({ queryKey, pageParam }) =>
      drugClient.paginated(Object.assign({}, queryKey[1], pageParam)),
    {
      keepPreviousData: true,
    },
  );

  return {
    drugs: data?.data ?? [],
    paginatorInfo: mapPaginatorData(data),
    error,
    loading: isLoading,
  };
};

export const useDrugQuery = ({ slug, language }: GetParams) => {
  const { data, error, isLoading } = useQuery<Drug, Error>(
    [API_ENDPOINTS.CATEGORIES, { slug, language }],
    () => drugClient.get({ slug, language }),
  );

  return {
    drug: data,
    error,
    isLoading,
  };
};

export const useCreateDrugMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(drugClient.create, {
    onSuccess: () => {
      Router.push(Routes.drug.list, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-created'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.DRUGS);
    },
    onError: (error: any) => {
      toast.error(error?.response?.data?.error ?? "Something went wrong")
    }
  });
};

export const useUpdateDrugMutation = () => {
  const { t } = useTranslation();
  const router = useRouter();
  const queryClient = useQueryClient();
  return useMutation(drugClient.update, {
    onSuccess: async (data) => {
      const generateRedirectUrl = router.query.shop
        ? `/${router.query.shop}${Routes.drug.list}`
        : Routes.drug.list;
      await router.push(`${generateRedirectUrl}/${data?.id}/edit`, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-updated'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.DRUGS);
    },
  });
};

export const useDeleteDrugMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(drugClient.delete, {
    onSuccess: () => {
      toast.success(t('common:successfully-deleted'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.DRUGS);
    },
  });
};
