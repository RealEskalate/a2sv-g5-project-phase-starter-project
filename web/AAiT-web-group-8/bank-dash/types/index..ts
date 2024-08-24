export interface SidebarProps {
    isOpen: boolean;
    toggleSidebar: () => void;
  }
  export interface CardProps {
    balance: string;
    cardHolder: string;
    validThru: string;
    cardNumber: string;
    gradientFrom: string;
    gradientTo: string;
    chipImage: string; 
    borderStyle?: string;
    bottomBackground: string;
    textColor : string;
  }
    
  export interface RecentTransactionProps {
    title: string;
    date: string;
    amount: string;
    type: 'income' | 'expense';
    imageSrc: string;
    }
  export interface LabelProps {
    cx: number;
    cy: number;
    midAngle: number;
    innerRadius: number;
    outerRadius: number;
    percent: number;
    index: number;
  }
  
  export interface PersonProps {
    person: {
      name: string;
      role: string;
      img: string;
    };
    selectedPerson: string | null;
  }
  
  
  export interface EditProfileProps {
    isActive: boolean;
  }
  
  export type FormValues = {
    name: string;
    email: string;
    dateOfBirth: string;
    permanentAddress: string;
    postalCode: string;
    username: string;
    password: string;
    presentAddress: string;
    city: string;
    country: string;
  };