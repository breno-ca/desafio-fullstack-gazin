import { NivelForm } from './nivel.model';

export type CreateNivelRequest = Omit<NivelForm, 'id'>;

export type UpdateNivelRequest = NivelForm;
