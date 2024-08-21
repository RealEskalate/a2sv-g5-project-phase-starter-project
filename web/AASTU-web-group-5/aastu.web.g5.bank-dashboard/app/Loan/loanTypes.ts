// types.ts
export interface LoanDataProps {
	id: string; // or number, depending on your data
	loanAmount: number;
	leftToRepay: number;
	duration: string; // Adjust based on the actual type
	interestRate: number;
	installment: number;
	type: string;
}
