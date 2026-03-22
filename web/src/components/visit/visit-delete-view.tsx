import ConfirmationCard from '@/components/common/confirmation-card';
import {
  useModalAction,
  useModalState,
} from '@/components/ui/modal/modal.context';
import { useDeletePatientMutation } from '@/data/patient';
import { getErrorMessage } from '@/utils/form-error';

const PatientDeleteView = () => {
  const { mutate: deletePatient, isLoading: loading } =
    useDeletePatientMutation();
  const { data } = useModalState();
  const { closeModal } = useModalAction();

  async function handleDelete() {
    try {
      deletePatient({ id: data });
      closeModal();
    } catch (error) {
      closeModal();
      getErrorMessage(error);
    }
  }

  return (
    <ConfirmationCard
      onCancel={closeModal}
      onDelete={handleDelete}
      deleteBtnLoading={loading}
    />
  );
};

export default PatientDeleteView;
