import { Desenvolvedor, DesenvolvedorForm } from './desenvolvedor.model';

export type CreateDesenvolvedorRequest = Omit<Desenvolvedor, 'id' | 'nivel'>;

export type UpdateDesenvolvedorRequest = DesenvolvedorForm;
