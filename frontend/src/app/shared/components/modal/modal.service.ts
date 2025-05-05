import { Injectable } from '@angular/core';
import { ModalComponent } from './modal.component';

@Injectable({
  providedIn: 'root',
})
export class ModalService {
  message: string | null = null;
  private resolver: ((result: boolean) => void) | null = null;

  confirm(message: string): Promise<boolean> {
    this.message = message;
    return new Promise<boolean>((resolve) => {
      this.resolver = resolve;
    });
  }

  resolve(result: boolean) {
    if (this.resolver) {
      this.resolver(result);
      this.reset();
    }
  }

  private reset() {
    this.message = null;
    this.resolver = null;
  }
}
