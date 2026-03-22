import * as yup from 'yup';

export const visitValidationSchema = yup.object().shape({
  patient: yup.object().required('form:error-name-required'),
  prompt: yup.string().required('form:error-prompt-required'),
});
