import ConfirmationCard from '@/components/common/confirmation-card';
import {
  useModalAction,
  useModalState,
} from '@/components/ui/modal/modal.context';
import { useDeleteLabTestMutation } from '@/data/labTest';
import { getErrorMessage } from '@/utils/form-error';

const LabTestDeleteView = () => {
  const { mutate: deleteLabTest, isLoading: loading } =
    useDeleteLabTestMutation();
  const { data } = useModalState();
  const { closeModal } = useModalAction();

  async function handleDelete() {
    try {
      deleteLabTest({ id: data });
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

export default LabTestDeleteView;
