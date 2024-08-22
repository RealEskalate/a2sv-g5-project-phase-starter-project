import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

const invoices = [
  {
    Description: "Spotify",
    Transaction: "#12312312",
    Type: "Etertainment",
    card: "121231****",
    Date: "28 Jan, 12:30AM",
    Amount: "-$12000",
  },
  {
    Description: "Spotify",
    Transaction: "#12312312",
    Type: "Etertainment",
    card: "121231****",
    Date: "28 Jan, 12:30AM",
    Amount: "-$12000",
  },
  {
    Description: "Spotify",
    Transaction: "#12312312",
    Type: "Etertainment",
    card: "121231****",
    Date: "28 Jan, 12:30AM",
    Amount: "-$12000",
  },
  {
    Description: "Spotify",
    Transaction: "#12312312",
    Type: "Etertainment",
    card: "121231****",
    Date: "28 Jan, 12:30AM",
    Amount: "-$12000",
  },
  {
    Description: "Spotify",
    Transaction: "#12312312",
    Type: "Etertainment",
    card: "121231****",
    Date: "28 Jan, 12:30AM",
    Amount: "-$12000",
  },
];

export function TableDemo() {
  return (
    <Table className="mb-10">
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">Description</TableHead>
          <TableHead>Transaction ID</TableHead>
          <TableHead>Type</TableHead>
          <TableHead>Card</TableHead>
          <TableHead>Date</TableHead>
          <TableHead>Amount</TableHead>
          <TableHead className="">Receipt</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {invoices.map((invoice) => (
          <TableRow key={invoice.Transaction}>
            <TableCell className="font-medium">{invoice.Description}</TableCell>
            <TableCell>{invoice.Transaction}</TableCell>
            <TableCell>{invoice.Type}</TableCell>
            <TableCell>{invoice.card}</TableCell>
            <TableCell>{invoice.Date}</TableCell>
            <TableCell className="text-right text-red-600">
              {invoice.Amount}
            </TableCell>
            <TableCell>
              <button className="rounded-3xl border-blue-600 p-2 text-blue-600 border-2 flex items-center justify-center">
                Download
              </button>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
      <TableFooter>
        <TableRow>
          <TableCell colSpan={5}>Total</TableCell>
          <TableCell className="text-right">$2,500.00</TableCell>
        </TableRow>
      </TableFooter>
    </Table>
  );
}
