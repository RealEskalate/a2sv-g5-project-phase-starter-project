import Image from "next/image";

interface prop {
  image: string;
  title: string;
  date: string;
  expense: number;
  color: string;
}

const Invoices = ({ image, title, date, expense, color }: prop) => {
  return (
    <div className="flex justify-between max-w-screen-sm min-w-[325px]">
      <div className="flex flex-initial w-[3/12] m-3">
        <div
          className={`${color} bg-opacity-25 font-semibold py-1 px-2 rounded-lg  text-sm w-[45px]`}
        >
          <Image src={`${image}`} alt={title} width={27} height={18} />
        </div>
        <div className="flex flex-col ">
          <div>{title}</div>
          <div className="text-[#718EBF]">{date}</div>
        </div>
      </div>
      <div className="flex-initial w-[4/12] m-3 text-[#718EBF]">${expense}</div>
    </div>
  );
};

export default Invoices;
