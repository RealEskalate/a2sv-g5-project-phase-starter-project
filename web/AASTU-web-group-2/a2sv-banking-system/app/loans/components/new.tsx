import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";

interface invoices {
  serialNumber: string;
  loanAmount: number;
  amountLeftToRepay: number;
  duration: number;
  interestRate: number;
  installment: number;
  type: string;
  activeLoneStatus: string;
  userId: string;
}

export function New({ invoice }: { invoice : invoices }) {
  return (
    <Sheet>
      <SheetTrigger asChild>
        <button className="">Details</button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Loan details</SheetTitle>
          <SheetDescription>
            if there are any mistakes you can contact the support team.
          </SheetDescription>
        </SheetHeader>
        <div className="pt-4">serialNumber      : {invoice.serialNumber}</div>
        <div>loanAmount        : {invoice.loanAmount}</div>
        <div>amountLeftToRepay : {invoice.amountLeftToRepay}</div>
        <div>duration          : {invoice.duration}</div>
        <div>interestRate      : {invoice.interestRate}</div>
        <div>installment       : {invoice.installment}</div>
        <div>type              : {invoice.type}</div>
        <div>userId            : {invoice.userId}</div>
      </SheetContent>
    </Sheet>
  );
}
