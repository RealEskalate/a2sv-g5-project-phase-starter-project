import React from 'react'
import InformationCard from './components/InformationCard'
import BankServiceList from './components/BankServiceList'
const page = () => {
  return (
    <div className='flex-col bg-[#f5f7fa]'>
      <div className='flex mx-5 my-4 rounded-3xl gap-4 overflow-x-auto [&::-webkit-scrollbar]:hidden'>
      <InformationCard
        logoBgColor="#e7edff"
        logo={
              <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <g clip-path="url(#clip0_163_357)">
              <path d="M18.2112 2.72841C15.6138 2.72841 13.6258 2.07743 12.4154 1.53132C11.0997 0.937663 10.3902 0.335245 10.3843 0.330201L10.001 0L9.61716 0.329224C9.61016 0.335245 8.90061 0.937702 7.58493 1.53132C6.37462 2.07743 4.38661 2.72841 1.78914 2.72841H1.20264V10.7793L10.0002 20L18.7977 10.7874V2.72841H18.2112ZM16.4517 8.20503C16.4517 9.11454 16.0983 9.9697 15.4564 10.6137L10.0002 16.2817L4.54395 10.6137C3.90204 9.96966 3.54865 9.1145 3.54865 8.20499V7.90008C3.54865 6.01862 5.07935 4.48788 6.96085 4.48788C7.80522 4.48788 8.61588 4.79897 9.24348 5.36381L9.97892 6.0257L10.5173 5.48729C11.1618 4.8428 12.0187 4.48785 12.9301 4.48785H13.0395C14.921 4.48785 16.4517 6.01854 16.4517 7.90004V8.20503H16.4517Z" fill="#396AFF"/>
              <path d="M13.0395 5.66093H12.9302C12.3321 5.66093 11.7698 5.89385 11.3468 6.31676L10.0215 7.6421L8.45884 6.2357C8.04696 5.86503 7.515 5.66089 6.96087 5.66089C5.72616 5.66089 4.72168 6.66537 4.72168 7.90008V8.20498C4.72168 8.8031 4.9546 9.36544 5.37751 9.78835L5.38533 9.79628L10.0002 14.5902L14.6229 9.78831C15.0458 9.3654 15.2787 8.80306 15.2787 8.20494V7.90004C15.2787 6.66545 14.2742 5.66093 13.0395 5.66093ZM11.7597 10.9395H10.5867V12.1125H9.41371V10.9395H8.2407V9.76645H9.41371V8.59344H10.5867V9.76645H11.7597V10.9395Z" fill="#396AFF"/>
              </g>
              <defs>
              <clipPath id="clip0_163_357">
              <rect width="20" height="20" fill="white"/>
              </clipPath>
              </defs>
              </svg>
        }
        title="Life Insurance"
        description="Unlimited protection"
        cardBgColor="bg-[#ffffff]"
      />
            <InformationCard
        logoBgColor="#fff5d9"
        logo={
<svg width="16" height="20" viewBox="0 0 16 20" fill="none" xmlns="http://www.w3.org/2000/svg">
<path d="M15.3125 17.275L14.2463 5.56875C14.2175 5.24625 13.9475 5 13.6238 5H11.7488V3.75C11.7488 2.745 11.3588 1.8025 10.6525 1.09625C9.95752 0.4 8.99002 0 7.99877 0C5.93127 0 4.24877 1.6825 4.24877 3.75V5H2.37377C2.05002 5 1.78002 5.24625 1.75127 5.56875L0.687524 17.2738C0.623774 17.9725 0.858775 18.6688 1.33127 19.1863C1.80377 19.7038 2.47627 20 3.17752 20H12.8213C13.5213 20 14.1938 19.7038 14.6663 19.1875C15.14 18.67 15.3738 17.9725 15.3125 17.275ZM10.4988 5H5.49877V3.75C5.49877 2.37125 6.62002 1.25 7.99877 1.25C8.66127 1.25 9.30627 1.515 9.76877 1.97875C10.24 2.45 10.4988 3.07875 10.4988 3.75V5Z" fill="#FFBB38"/>
</svg>

        }
        title="Shopping"
        description="Buy. Think. Grow"
        cardBgColor="bg-[#ffffff]"
      />
            <InformationCard
        logoBgColor="#dcfaf8"
        logo={
          <svg width="18" height="20" viewBox="0 0 18 20" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M16.6553 13.3559C16.1178 14.8128 15.3048 16.0796 14.2387 17.1207C13.0251 18.3057 11.4361 19.2473 9.5156 19.919C9.45258 19.941 9.38681 19.959 9.32074 19.9722C9.23346 19.9895 9.14466 19.9988 9.05661 20H9.03937C8.94553 20 8.85123 19.9905 8.75769 19.9722C8.69162 19.959 8.62677 19.941 8.56406 19.9194C6.6413 19.2488 5.05026 18.3077 3.83551 17.1227C2.76892 16.0815 1.95608 14.8155 1.41928 13.3586C0.443171 10.7097 0.498713 7.79159 0.543421 5.44662L0.544184 5.41061C0.553187 5.21697 0.558985 5.01357 0.562189 4.78896C0.578516 3.68621 1.45529 2.77389 2.55819 2.71239C4.85769 2.58407 6.63657 1.8341 8.1565 0.35262L8.16977 0.340413C8.42215 0.108937 8.74014 -0.00458834 9.05661 0.00014191C9.36179 0.00410921 9.66574 0.117482 9.90912 0.340413L9.92209 0.35262C11.4423 1.8341 13.2212 2.58407 15.5207 2.71239C16.6236 2.77389 17.5004 3.68621 17.5167 4.78896C17.5199 5.0151 17.5257 5.21819 17.5347 5.41061L17.5352 5.42587C17.5797 7.77526 17.635 10.6992 16.6553 13.3559Z" fill="#16DBCC"/>
          <path d="M16.6554 13.356C16.1178 14.8129 15.3048 16.0797 14.2387 17.1208C13.0251 18.3058 11.4361 19.2474 9.51562 19.9191C9.45261 19.9411 9.38684 19.9591 9.32077 19.9723C9.23349 19.9896 9.14468 19.9989 9.05664 20.0001V0.000244141C9.36182 0.00421144 9.66577 0.117584 9.90915 0.340515L9.92212 0.352722C11.4424 1.8342 13.2212 2.58417 15.5207 2.7125C16.6236 2.77399 17.5004 3.68631 17.5167 4.78907C17.5199 5.0152 17.5257 5.2183 17.5347 5.41071L17.5352 5.42597C17.5797 7.77537 17.635 10.6993 16.6554 13.356Z" fill="#16DBCC"/>
          <path d="M14.0237 10.0001C14.0237 12.7425 11.797 14.9749 9.05682 14.9847H9.03927C6.29101 14.9847 4.05469 12.7485 4.05469 10.0001C4.05469 7.2518 6.29101 5.01562 9.03927 5.01562H9.05682C11.797 5.02539 14.0237 7.25775 14.0237 10.0001Z" fill="white"/>
          <path d="M12.1015 9.19753L9.07623 11.8814L8.42252 12.4613C8.26809 12.5983 8.06554 12.6667 7.86319 12.6667C7.66064 12.6667 7.4583 12.5983 7.30366 12.4613L5.89815 11.214C5.58929 10.94 5.58929 10.4962 5.89815 10.2221C6.2066 9.94812 6.70753 9.94812 7.01639 10.2221L7.86319 10.9732L10.9833 8.20561C11.2921 7.93146 11.7931 7.93146 12.1015 8.20561C12.4104 8.47959 12.4104 8.92392 12.1015 9.19753Z" fill="#16DBCC"/>
          </svg>
        }
        title="Safety"
        description="We are your allies"
        cardBgColor="bg-[#ffffff]"
      />
      
      
      </div>
      <h1 className='text-[#343C6A] font-semibold mx-5 my-4 text-xl md:font-bold'>Bank Services List</h1>
        <div className='flex-col gap-5'>
          <BankServiceList
            logoBgColor="bg-[#FFE0EB]"
            logoSvg={(
              <svg className="w-8 h-8 text-blue-500 items-center" fill="currentColor" viewBox="0 0 24 24">
                <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <g clip-path="url(#clip0_163_431)">
                    <path d="M11.5941 12.2698C14.95 12.2698 17.6803 9.51765 17.6803 6.1349C17.6803 2.75215 14.95 0 11.5941 0C8.2381 0 5.50781 2.75211 5.50781 6.13486C5.50781 9.51761 8.2381 12.2698 11.5941 12.2698ZM9.63836 8.04927C9.8154 7.77865 10.1784 7.70268 10.449 7.8798C10.8433 8.13773 10.9921 8.16151 11.5124 8.15788C12.0203 8.15452 12.315 7.77595 12.3739 7.42562C12.4026 7.25521 12.4135 6.83909 11.8978 6.65678C11.2929 6.44292 10.6738 6.1985 10.2432 5.86069C9.81248 5.52288 9.61525 4.93971 9.7285 4.33886C9.85128 3.68749 10.3057 3.16897 10.9144 2.98564C10.9199 2.984 10.9253 2.98268 10.9308 2.98104V2.75906C10.9308 2.43566 11.1929 2.17347 11.5163 2.17347C11.8397 2.17347 12.1019 2.43566 12.1019 2.75906V2.94414C12.4997 3.03909 12.7774 3.22105 12.8902 3.30534C13.1492 3.49901 13.2023 3.86594 13.0086 4.125C12.815 4.38407 12.4481 4.43708 12.189 4.24337C12.069 4.15366 11.7373 3.9608 11.2521 4.10701C10.9687 4.19242 10.8952 4.4721 10.8794 4.55576C10.8484 4.72016 10.8832 4.87425 10.9659 4.93913C11.2645 5.17329 11.8019 5.38062 12.2881 5.55251C13.1848 5.86947 13.6835 6.70023 13.5289 7.61976C13.453 8.07094 13.2261 8.48951 12.8897 8.79851C12.6607 9.00901 12.394 9.15908 12.1019 9.24477V9.51066C12.1019 9.83406 11.8397 10.0963 11.5163 10.0963C11.1929 10.0963 10.9308 9.83406 10.9308 9.51066V9.30317C10.5521 9.25726 10.2343 9.13885 9.80779 8.85988C9.53721 8.68284 9.46132 8.31989 9.63836 8.04927Z" fill="#FF82AC"/>
                    <path d="M2.21964 14.2373H0.884417C0.561016 14.2373 0.298828 14.4995 0.298828 14.8229V19.4135C0.298828 19.7369 0.561016 19.9991 0.884417 19.9991H2.21968V14.2373H2.21964Z" fill="#FF82AC"/>
                    <path d="M19.5295 14.1965C18.4319 13.0989 16.646 13.0989 15.5485 14.1965L13.7943 15.9507L13.0753 16.6697C12.7847 16.9602 12.3906 17.1235 11.9797 17.1235H8.48353C8.16778 17.1235 7.89607 16.8808 7.8812 16.5654C7.86535 16.2287 8.13366 15.9507 8.46694 15.9507H12.0205C12.735 15.9507 13.3547 15.442 13.4776 14.7382C13.5058 14.5765 13.5205 14.4104 13.5205 14.2408C13.5205 13.9168 13.258 13.6539 12.9341 13.6539H10.9869C10.3506 13.6539 9.7395 13.3652 9.0925 13.0596C8.41389 12.739 7.71219 12.4075 6.89171 12.3529C6.17409 12.3051 5.45483 12.3837 4.7538 12.5861C4.00319 12.8029 3.46363 13.4697 3.3982 14.2397C3.3957 14.2395 3.39316 14.2394 3.39062 14.2393V19.9972L13.4794 20C14.1731 20 14.8253 19.7298 15.3158 19.2393L19.5293 15.0258C19.7585 14.7969 19.7585 14.4255 19.5295 14.1965Z" fill="#FF82AC"/>
                    </g>
                    <defs>
                    <clipPath id="clip0_163_431">
                    <rect width="20" height="20" fill="white"/>
                    </clipPath>
                  </defs>
                  </svg>

              </svg>
            )}
            serviceName="Custom Service"
            serviceDescription="This is a custom description"
            additionalServices={[
              { name: "Service 1", description: "Service 1 Description" },
              { name: "Service 2", description: "Service 2 Description" },
              { name: "Service 3", description: "Service 3 Description" },
            ]}
            viewDetailsLink="https://example.com/details"
          />
                    <BankServiceList
            logoBgColor="bg-[#FFF5D9]"
            logoSvg={(
              <svg className="w-8 h-8 text-blue-500" fill="currentColor" viewBox="0 0 24 24">
                {/* Custom SVG content */}
              </svg>
            )}
            serviceName="Custom Service"
            serviceDescription="This is a custom description"
            additionalServices={[
              { name: "Service 1", description: "Service 1 Description" },
              { name: "Service 2", description: "Service 2 Description" },
              { name: "Service 3", description: "Service 3 Description" },
            ]}
            viewDetailsLink="https://example.com/details"
          />
                    <BankServiceList
            logoBgColor="bg-[#FFE0EB]"
            logoSvg={(
              <svg className="w-8 h-8 text-blue-500" fill="currentColor" viewBox="0 0 24 24">
                {/* Custom SVG content */}
              </svg>
            )}
            serviceName="Custom Service"
            serviceDescription="This is a custom description"
            additionalServices={[
              { name: "Service 1", description: "Service 1 Description" },
              { name: "Service 2", description: "Service 2 Description" },
              { name: "Service 3", description: "Service 3 Description" },
            ]}
            viewDetailsLink="https://example.com/details"
          />
                    <BankServiceList
            logoBgColor="bg-[#E7EDFF]"
            logoSvg={(
              <svg className="w-8 h-8 text-blue-500" fill="currentColor" viewBox="0 0 24 24">
                {/* Custom SVG content */}
              </svg>
            )}
            serviceName="Custom Service"
            serviceDescription="This is a custom description"
            additionalServices={[
              { name: "Service 1", description: "Service 1 Description" },
              { name: "Service 2", description: "Service 2 Description" },
              { name: "Service 3", description: "Service 3 Description" },
            ]}
            viewDetailsLink="https://example.com/details"
          />
                    <BankServiceList
            logoBgColor="bg-[#DCFAF8]"
            logoSvg={(
              <svg className="w-8 h-8 text-blue-500" fill="currentColor" viewBox="0 0 24 24">
                {/* Custom SVG content */}
              </svg>
            )}
            serviceName="Custom Service"
            serviceDescription="This is a custom description"
            additionalServices={[
              { name: "Service 1", description: "Service 1 Description" },
              { name: "Service 2", description: "Service 2 Description" },
              { name: "Service 3", description: "Service 3 Description" },
            ]}
            viewDetailsLink="https://example.com/details"
          />
                    <BankServiceList
            logoBgColor="bg-[#FFE0EB]"
            logoSvg={(
              <svg className="w-8 h-8 text-blue-500" fill="currentColor" viewBox="0 0 24 24">
                {/* Custom SVG content */}
              </svg>
            )}
            serviceName="Custom Service"
            serviceDescription="This is a custom description"
            additionalServices={[
              { name: "Service 1", description: "Service 1 Description" },
              { name: "Service 2", description: "Service 2 Description" },
              { name: "Service 3", description: "Service 3 Description" },
            ]}
            viewDetailsLink="https://example.com/details"
          />
      </div>



    </div>
  )
}

export default page
