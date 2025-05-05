import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { DesenvolvedorForm } from '../models/desenvolvedor.model';
import { NivelSelectOptions } from '../../niveis/models/nivel.model';
import { NiveisService } from '../../niveis/services/niveis.service';
import { ToastService } from 'src/app/shared/components/toast/toast.service';
import { FormField } from 'src/app/shared/components/form/form.model';

@Component({
  selector: 'app-desenvolvedores-form',
  templateUrl: './desenvolvedores-form.component.html',
  styleUrls: ['./desenvolvedores-form.component.scss'],
})
export class DesenvolvedoresFormComponent implements OnInit {
  @Input() initialData: DesenvolvedorForm | null = null;
  @Output() close = new EventEmitter<DesenvolvedorForm | null>();

  niveis: NivelSelectOptions[] = [];

  form = this.fb.group({
    id: [''],
    nivel_id: [''],
    nome: [''],
    sexo: [''],
    data_nascimento: [''],
    hobby: [''],
  });

  fields: FormField[] = [
    { label: 'ID', name: 'id', type: 'text', hidden: true, disabled: true },
    {
      label: 'Nome',
      name: 'nome',
      type: 'text',
      disabled: false,
      validators: [Validators.required, Validators.min(3), Validators.max(100)],
      invalidMessage: 'Obrigatório preencher o nome.',
    },
    {
      label: 'Nível',
      name: 'nivel_id',
      type: 'select',
      disabled: false,
      fullWidth: false,
      options: this.niveis.map((n) => ({ label: n.nivel, value: n.id })),
      invalidMessage: 'Obrigatório preencher o nível.',
    },
    {
      label: 'Sexo',
      name: 'sexo',
      type: 'select',
      disabled: false,
      fullWidth: false,
      options: [
        { label: 'M', value: 'M' },
        { label: 'F', value: 'F' },
      ],
      validators: [Validators.required],
      invalidMessage: 'Obrigatório preencher o sexo.',
    },
    {
      label: 'Data de Nascimento',
      name: 'data_nascimento',
      type: 'date',
      disabled: false,
      validators: [Validators.required],
      invalidMessage: 'Obrigatório preencher a data de nascimento.',
    },
    {
      label: 'Hobby',
      name: 'hobby',
      type: 'text',
      disabled: false,
      validators: [Validators.required, Validators.min(3), Validators.max(100)],
      invalidMessage: 'Obrigatório preencher o Hobby.',
    },
  ];

  get idEdit(): boolean {
    return !!this.initialData?.id;
  }

  constructor(
    private fb: FormBuilder,
    private niveisService: NiveisService,
    private toastService: ToastService,
  ) {}

  ngOnInit(): void {
    this.niveisService.getSelectOptions().subscribe({
      next: (response: NivelSelectOptions[]) => {
        this.niveis = response;

        this.fields[2] = {
          label: 'Nível',
          name: 'nivel_id',
          type: 'select',
          disabled: false,
          options: this.niveis.map((n) => ({ label: n.nivel, value: n.id })),
        };
      },
      error: () => {
        this.toastService.show('Não há niveis para listar', true);
      },
    });

    if (this.initialData) {
      this.form.patchValue(this.initialData);
    }
  }
  onSubmit(desenvolvedor: DesenvolvedorForm): void {
    this.close.emit(desenvolvedor);
  }

  onCancel(): void {
    this.close.emit(null);
  }
}
