import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ToastService {
  message: string | null = null;
  isError = false;

  show(message: string, isError = false) {
    this.message = message;
    this.isError = isError;

    setTimeout(() => {
      this.message = null;
    }, 3500);
  }
}
