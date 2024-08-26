import { useForm } from "react-hook-form";
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
import Card1 from "./Card1";
type Form = {
  loanAmount: number;
  duration: number;
  interestRate: number;
  type: string
}

export function SheetDemo({ handleform }: { handleform : (data:Form)=>void}) {
  const forms = useForm<Form>();
  const { register, control, handleSubmit, formState, reset } = forms;
  const { errors } = formState;
  const onSubmit = (data: Form) => {
    handleform(data);
    console.log("formsubmited", data);
    reset();
  };
  return (
    <Sheet>
      <SheetTrigger asChild>
        <button className="border-l-2 pl-2 border-blue-200">
          <Card1 text="Choose Loans" img="/custom.svg" num="Loan Money" />
        </button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle>Loan Money</SheetTitle>
          <SheetDescription>
            You can choose any type of your choice.
          </SheetDescription>
        </SheetHeader>
        <form onSubmit={handleSubmit(onSubmit)} noValidate className="mt-3">
          <label className="loanAmount" htmlFor="loanAmount ">
            {" "}
            Loan Amount{" "}
          </label>
          <input
            type="text"
            id="loanAmount"
            {...register("loanAmount", {
              required: "loan Amount is requered",
              pattern: {
                value: /^[0-9]+$/,
                message: "invalid Loan Amount",
              },
            })}
            className="border w-full border-[#DFEAF2] focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"

            // className={errors.loanAmount && "error"}
          />{" "}
          <p className="text-red-700">{errors.loanAmount?.message}</p>
          <div className="mt-5">
            <label className="" htmlFor="duration">
              {" "}
              Duration{" "}
            </label>
          </div>
          <input
            type="text"
            id="duration"
            {...register("duration", {
              required: "duration is requered",
              pattern: {
                value: /^[0-9]+$/,
                message: "invalid duration",
              },
            })}
            // className={errors.duration && "error"}
            className="border w-full border-[#DFEAF2] focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
          />
          <p className="text-red-700">{errors.duration?.message}</p>
          <div className="mt-5">
            <label className="interestRate" htmlFor="interestRate">
              {" "}
              Interest Rate{" "}
            </label>
          </div>
          <input
            type="text"
            id="interestRate "
            {...register("interestRate", {
              required: "interestRate is requered",
              pattern: {
                value: /^[0-9]+$/,
                message: "invalid interestRate",
              },
            })}
            // className={errors.interestRate && "error"}
            className="border w-full border-[#DFEAF2] focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
          />
          <p className="text-red-700">{errors.interestRate?.message}</p>
          <div className="mt-5">
            <label className="type" htmlFor="type">
              {" "}
              Type{" "}
            </label>
          </div>
          <select
            id="dropdown"
            // name="dropdown"
            {...register("type", {
              required: "type is requered",
            })}
            className="border w-full border-[#DFEAF2] focus:outline-none focus:ring-2 focus:ring-blue-500 rounded-xl py-3 px-6 placeholder:text-[#718EBF] dark:border-gray-600 dark:focus:outline-none dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
          >
            <option value="personal">Personal</option>
            <option value="corporate">Corporate</option>
            <option value="business">Business</option>
          </select>
          <br />
          <p className="text-red-700">{errors.type?.message}</p>
          <div className="flex justify-center">
            <button className="bg-blue-500  text-white font-semibold py-2 px-4 rounded shadow hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400 mt-5">
              Submit
            </button>
          </div>
        </form>
      </SheetContent>
    </Sheet>
  );
}
