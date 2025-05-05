import { ValidatorFn } from '@angular/forms';

export interface FormField {
  label: string;
  name: string;
  type?: 'text' | 'date' | 'number' | 'select';
  options?: { label: string; value: any }[];
  hidden?: boolean;
  disabled: boolean;
  fullWidth?: boolean;
  validators?: ValidatorFn[];
  invalidMessage?: string;
}
