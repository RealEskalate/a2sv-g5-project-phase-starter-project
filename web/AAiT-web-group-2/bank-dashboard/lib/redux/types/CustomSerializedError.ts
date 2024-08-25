import { SerializedError } from '@reduxjs/toolkit';

export interface CustomSerializedError extends SerializedError {
  data: {
    message: string;
  };
}