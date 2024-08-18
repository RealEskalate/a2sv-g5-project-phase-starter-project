import * as React from "react"
const SafetyIcon = (props) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={50}
    height={50}
    fill="none"
    {...props}
  >
    <circle cx={25} cy={25} r={25} fill="#DCFAF8" />
    <g clipPath="url(#a)">
      <path
        fill="#16DBCC"
        d="M32.655 28.356a10.09 10.09 0 0 1-2.416 3.765c-1.214 1.185-2.803 2.126-4.723 2.798a1.464 1.464 0 0 1-.46.081h-.017a1.467 1.467 0 0 1-.475-.08c-1.923-.671-3.514-1.612-4.729-2.797a10.072 10.072 0 0 1-2.416-3.764c-.976-2.65-.92-5.567-.876-7.912l.001-.036c.01-.194.015-.397.018-.622a2.12 2.12 0 0 1 1.996-2.077c2.3-.128 4.079-.878 5.599-2.36l.013-.012a1.28 1.28 0 0 1 1.74 0l.012.013c1.52 1.481 3.3 2.231 5.599 2.36a2.12 2.12 0 0 1 1.996 2.076c.003.226.009.43.018.622v.015c.045 2.35.1 5.273-.88 7.93Z"
      />
      <path
        fill="#16DBCC"
        d="M32.655 28.356a10.091 10.091 0 0 1-2.416 3.765c-1.214 1.185-2.803 2.126-4.723 2.798a1.467 1.467 0 0 1-.46.081V15c.306.004.61.118.853.34l.013.013c1.52 1.481 3.3 2.231 5.599 2.36a2.12 2.12 0 0 1 1.996 2.076c.003.226.009.43.018.622v.015c.045 2.35.1 5.273-.88 7.93Z"
      />
      <path
        fill="#fff"
        d="M30.024 25a4.99 4.99 0 0 1-4.967 4.985h-.018A4.99 4.99 0 0 1 20.055 25a4.99 4.99 0 0 1 4.984-4.984h.018A4.99 4.99 0 0 1 30.024 25Z"
      />
      <path
        fill="#16DBCC"
        d="m28.102 24.198-3.026 2.683-.654.58a.843.843 0 0 1-.559.206.843.843 0 0 1-.56-.206l-1.405-1.247a.647.647 0 0 1 0-.992.863.863 0 0 1 1.118 0l.847.751 3.12-2.767a.863.863 0 0 1 1.119 0 .647.647 0 0 1 0 .992Z"
      />
    </g>
    <defs>
      <clipPath id="a">
        <path fill="#fff" d="M15 15h20v20H15z" />
      </clipPath>
    </defs>
  </svg>
)
export default SafetyIcon
