import { Component, EventEmitter, Input, Output } from '@angular/core';
import { NivelForm } from '../models/nivel.model';
import { FormBuilder, Validators } from '@angular/forms';
import { FormField } from 'src/app/shared/components/form/form.model';

@Component({
  selector: 'app-niveis-form',
  templateUrl: './niveis-form.component.html',
  styleUrls: ['./niveis-form.component.scss'],
})
export class NiveisFormComponent {
  @Input() initialData: NivelForm | null = null;
  @Output() close = new EventEmitter<NivelForm | null>();

  form = this.fb.group({
    nivel: [''],
    id: [''],
  });

  fields: FormField[] = [
    { label: 'ID', name: 'id', hidden: true, disabled: true },
    {
      label: 'Nivel',
      name: 'nivel',
      fullWidth: true,
      disabled: false,
      validators: [Validators.required, Validators.min(3), Validators.max(100)],
      invalidMessage: 'Obrigatório preencher o nível.',
    },
  ];

  get idEdit(): boolean {
    return !!this.initialData?.id;
  }

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    if (this.initialData) {
      this.form.patchValue(this.initialData);
    }
  }

  onSubmit(nivel: NivelForm) {
    this.close.emit(nivel);
  }

  onCancel() {
    this.close.emit(null);
  }
}
