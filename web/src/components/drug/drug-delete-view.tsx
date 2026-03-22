import ConfirmationCard from '@/components/common/confirmation-card';
import {
  useModalAction,
  useModalState,
} from '@/components/ui/modal/modal.context';
import { useDeleteDrugMutation } from '@/data/drug';
import { getErrorMessage } from '@/utils/form-error';

const DrugDeleteView = () => {
  const { mutate: deleteDrug, isLoading: loading } =
    useDeleteDrugMutation();
  const { data } = useModalState();
  const { closeModal } = useModalAction();

  async function handleDelete() {
    try {
      deleteDrug({ id: data });
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

export default DrugDeleteView;
