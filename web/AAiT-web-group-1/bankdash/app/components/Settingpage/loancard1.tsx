'use client';
import icon1 from '../../../public/icon1.svg'
import icon2 from '../../../public/icon2.svg'
import icon3 from '../../../public/icon3.svg'
import icon4 from '../../../public/icon4.svg'
const LoanCard = () => {
    return (
        <div className="flex flex-row gap-4">
            <div className="flex flex-row items-center w-64 h-32 rounded-2xl border p-2 shadow-lg">
                <div className="mr-4">
                    <img src={icon1.src} alt="" />
                </div>
                <div>
                    <p className="text-base" style={{ color: '#718EBF' }}>Personal Loans</p>
                    <h3 className="font-bold">$50,000</h3>
                </div>
            </div>

            <div className="flex flex-row items-center w-64 h-32 rounded-2xl border p-2 shadow-lg">
                <div className="mr-4">
                    <img src={icon2.src} alt="" />
                </div>
                <div>
                    <p className="text-base" style={{ color: '#718EBF' }}>Corprate Loans</p>
                    <h3 className="font-bold">$50,000</h3>
                </div>
            </div>

            <div className="flex flex-row items-center w-64 h-32 rounded-2xl border p-2 shadow-lg">
                <div className="mr-4">
                    <img src={icon3.src} alt="" />
                </div>
                <div>
                    <p className="text-base" style={{ color: '#718EBF' }}>Business Loans</p>
                    <h3 className="font-bold">$50,000</h3>
                </div>
            </div>

            <div className="flex flex-row items-center w-64 h-32 rounded-2xl border p-2 shadow-lg">
                <div className="mr-4">
                    <img src={icon4.src} alt="" />
                </div>
                <div>
                    <p className="text-base" style={{ color: '#718EBF' }}>Custom Loans</p>
                    <h3 className="font-bold">Choose money</h3>
                </div>
            </div>
        </div>
    );
}

export default LoanCard;