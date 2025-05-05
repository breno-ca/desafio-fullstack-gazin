import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { FormField } from './form.model';

@Component({
  selector: 'app-form',
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.scss'],
  standalone: true,
})
export class FormComponent {
  @Input() fields: FormField[] = [];
  @Input() initialData: any = {};

  @Output() submit = new EventEmitter<any>();
  @Output() cancel = new EventEmitter<void>();

  form!: FormGroup;

  constructor(private fb: FormBuilder) {}

  ngOnInit() {
    const group: Record<string, any> = {};
    this.fields.forEach((field) => {
      group[field.name] = [this.initialData[field.name] || '', field.validators || []];
    });

    this.form = this.fb.group(group);
  }

  ngOnChanges() {
    if (this.form && this.initialData) {
      this.form.patchValue(this.initialData);
    }
  }

  onSubmit() {
    if (this.form.valid) {
      this.submit.emit(this.form.value);
    }
  }

  onCancel() {
    this.cancel.emit();
  }
}
