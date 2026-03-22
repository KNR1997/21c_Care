import Router, { useRouter } from 'next/router';
import { useMutation, useQuery, useQueryClient } from 'react-query';
import { toast } from 'react-toastify';
import { useTranslation } from 'next-i18next';
import { Routes } from '@/config/routes';
import { API_ENDPOINTS } from './client/api-endpoints';
import {
  Patient,
  PatientPaginator,
  PatientQueryOptions,
  GetParams,
} from '@/types';
import { mapPaginatorData } from '@/utils/data-mappers';
import { patientClient } from './client/patient';
import { Config } from '@/config';

export const usePatientsQuery = (options: Partial<PatientQueryOptions>) => {
  const { data, error, isLoading } = useQuery<PatientPaginator, Error>(
    [API_ENDPOINTS.PATIENTS, options],
    ({ queryKey, pageParam }) =>
      patientClient.paginated(Object.assign({}, queryKey[1], pageParam)),
    {
      keepPreviousData: true,
    },
  );

  return {
    patients: data?.data ?? [],
    paginatorInfo: mapPaginatorData(data),
    error,
    loading: isLoading,
  };
};

export const usePatientQuery = ({ slug, language }: GetParams) => {
  const { data, error, isLoading } = useQuery<Patient, Error>(
    [API_ENDPOINTS.CATEGORIES, { slug, language }],
    () => patientClient.get({ slug, language }),
  );

  return {
    patient: data,
    error,
    isLoading,
  };
};

export const useCreatePatientMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(patientClient.create, {
    onSuccess: () => {
      Router.push(Routes.patient.list, undefined, {
        locale: Config.defaultLanguage,
      });
      toast.success(t('common:successfully-created'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.PATIENTS);
    },
  });
};

export const useUpdatePatientMutation = () => {
  const { t } = useTranslation();
  const router = useRouter();
  const queryClient = useQueryClient();
  return useMutation(patientClient.update, {
    onSuccess: async (data) => {
      const generateRedirectUrl = router.query.shop
        ? `/${router.query.shop}${Routes.patient.list}`
        : Routes.patient.list;
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

export const useDeletePatientMutation = () => {
  const queryClient = useQueryClient();
  const { t } = useTranslation();

  return useMutation(patientClient.delete, {
    onSuccess: () => {
      toast.success(t('common:successfully-deleted'));
    },
    // Always refetch after error or success:
    onSettled: () => {
      queryClient.invalidateQueries(API_ENDPOINTS.PATIENTS);
    },
  });
};
