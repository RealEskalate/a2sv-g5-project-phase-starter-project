import { IconType } from "react-icons";

export interface Props {
  children: React.ReactNode;
}

export type ElementType = {
  text: string;
  destination: string;
  icon: IconType;
};

export interface navigationValue {
  activePage: string;
  toggle: boolean;
}
