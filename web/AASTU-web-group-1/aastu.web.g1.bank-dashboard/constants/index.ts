import { Description } from "@radix-ui/react-dialog"

export const sidebarLinks = [
    {
        title: 'Dashboard',
        route: '/dashboard',
        icon: '/icons/Home.svg'
    }, 
    {
        title: 'Transactions',
        route: '/dashboard/transactions',
        icon: '/icons/Transactions.svg'
    }, 
    {
        title: 'Accounts',
        route: '/dashboard/accounts',
        icon: '/icons/Accounts.svg'
    }, 

    {
        title: 'Investments',
        route: '/dashboard/investments',
        icon: '/icons/Investments.svg'
    }, 
    {
        title: 'Credit Cards',
        route: '/dashboard/credit-cards',
        icon: '/icons/CreditCard.svg'
    }, 
    {
        title: 'Loans',
        route: '/dashboard/loans',
        icon: '/icons/Loans.svg'
    }, 
    {
        title: 'Services',
        route: '/dashboard/services',
        icon: '/icons/Services.svg'
    }, 
    {
        title: 'My Privileges',
        route: '/dashboard/my-privileges',
        icon: '/icons/MyPrivileges.svg'
    }, 
    {
        title: 'Setting',
        route: '/dashboard/setting',
        icon: '/icons/Settings.svg'
    }, 
]

export const loanTypes =[
    {
        name: 'Personal Loans',
        route: "",
        icon: '/icons/PersonalLoan.svg',
        description:'50,000'
    },
    {
        name: 'Corporate Loans',
        route:'',
        icon: '/icons/CorporateLoan.svg',
        description:'100,000'
    },
    {
        name: 'Business Loans',
        route: '',
        icon: '/icons/BusinessLoan.svg',
        description:'500,000'
    },
    {
        name: 'Custom Loans',
        route: '',
        icon: '/icons/CustomLoan.svg',
        description:'50,000'
    },
]

export const investmentTypes =[
    {
        name: 'Total Invested Amount',
        route: "",
        icon: '/icons/Investment1.svg',
        description:'$150,000'
    },
    {
        name: 'Corporate Investments',
        route:'',
        icon: '/icons/Investment2.svg',
        description:'1250'
    },
    {
        name: 'Business Investments',
        route: '',
        icon: '/icons/Investment3.svg',
        description:'+5.80%'
    },
]

export const investmentsArray = [
    {
        name: 'Apple Store',
        type: ["E-commerce","Marketplace"],
        route: "",
        investmentValue: "$54,000",
        icon: '/icons/Apple.svg',
        returnValue:'+16%'
    },
    {
        name: 'Samsung Mobile',
        type: ["E-commerce","Marketplace"],
        route: "",
        investmentValue: '$25,300',
        icon:'/icons/Google.svg',
        returnValue:'-4%'
    },
    {
        name: 'Tesla Motors',
        type: ["Electric Vehicles"],
        route: "",
        investmentValue: '$8,200',
        icon:'/icons/Tesla.svg',
        returnValue:'+25%'
    },
]