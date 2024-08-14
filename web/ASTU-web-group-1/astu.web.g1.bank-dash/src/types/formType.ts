// Form Types - Input Group ( Label and INput Fiels )
export interface InputGroupType {
  id: string;
  label: string;
  inputType: string;
  registerName: string;
  register: any;
  placeholder: string;
}

export interface ToggleInputType {
  label: string;
  path: string;
  currentState: boolean;
}