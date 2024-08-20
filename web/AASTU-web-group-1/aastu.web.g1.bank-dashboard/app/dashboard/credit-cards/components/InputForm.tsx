const InputForm = () => {
    return (
      <div>
        <form>
          <div className="md:grid md:grid-cols-2 gap-4">
            <div className="flex flex-col my-4">
              <label className="text-[#515B6F] font-semibold">Card Type</label>
              <input
                type="text"
                placeholder="Classic"
                className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
              />
              {/* <p className="text-red-500 text-center mt-2">
                  {errors.name?.message}
                </p> */}
            </div>
            <div className="flex flex-col my-4">
              <label className="text-[#515B6F] font-semibold">Name On Card</label>
              <input
                type="text"
                placeholder="My Cards"
                className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
                // {...register("email")}
              />
              {/* <p className="text-red-500 text-center mt-2">
                  {errors.email?.message}
                </p> */}
            </div>
          </div>
          <div className="md:grid md:grid-cols-2 gap-4">
            <div className="flex flex-col my-4">
              <label className="text-[#515B6F] font-semibold">Card Number</label>
              <input
                type="number"
                placeholder="**** **** **** ****"
                className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
                // {...register("password")}
              />
              {/* <p className="text-red-500 text-center mt-2">
                  {errors.password?.message}
                </p> */}
            </div>
            <div className="flex flex-col my-4">
              <label className="text-[#515B6F] font-semibold">
                Expiration Date
              </label>
              <input
                type="date"
                placeholder="00:00:00 UTC on 1st January 1970"
                className="inputField mb-2 rounded-xl py-2 px-2 border border-gray-300"
                // {...register("confirmPassword")}
              />
              {/* <p className="text-red-500 text-center mt-2">
                  {errors.confirmPassword?.message}
                </p> */}
            </div>
          </div>
          {/* {error && (
              <p className="text-red-500 text-center mt-2 mb-5">{error}</p>
            )} */}
          <button className="bg-[#1814F3] sm:w-[100%] text-white p-2 sm:rounded-full md:max-w-[160px] md:rounded-md  ">
            Add to Cart
          </button>
        </form>
      </div>
    );
  };
export default InputForm  