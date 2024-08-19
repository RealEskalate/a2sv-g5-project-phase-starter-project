import Image from "next/image";

interface prop {
  image: string;
  alttext: string;
  description: string;
  transaction: number;
  type: string;
  account: string;
  status: string;
  colorimg: string;
  date: string;
}
const LastTransaction = ({
  image,
  alttext,
  description,
  transaction,
  colorimg,
  date,
  type,
  account,
  status,
}: prop) => {
  const color = transaction < 0 ? "text-red-500" : "text-green-500";

  // Format the transaction value with a "+" or "-" sign
  const formattedTransaction =
    transaction < 0 ? (
      <span className="text-red-500">-${Math.abs(transaction)}</span>
    ) : (
      <span className="text-green-500">+${transaction}</span>
    );

  return (
    <div>
      <div className="md:grid md:grid-cols-6 justify-between flex grid-cols-2 bg-white rounded-xl mb-2 mt-4 items-center min-w-[325px]">
        <div className="flex flex-initial col-span-2 m-3">
          <div
            className={`${colorimg} bg-opacity-25 font-semibold py-1 px-2 rounded-lg  text-sm w-[45px]`}
          >
            <Image src={`${image}`} alt={alttext} width={27} height={18} />
          </div>
          <div className="flex flex-col">
            <div>{description}</div>
            <div className="text-[#718EBF]">{date}</div>
          </div>
        </div>
        <div className="hidden md:block flex-initial w-[2/12]">{type}</div>
        <div className="hidden md:block flex-initial w-[2/12]">{account}</div>
        <div className="hidden md:block flex-initial w-[2/12]">{status}</div>
        <div className="flex-initial w-[4/12] m-3">{formattedTransaction}</div>
      </div>
    </div>
  );
};

export default LastTransaction;
