import { render, screen, waitFor } from "@testing-library/react";
import "@testing-library/jest-dom"; // Import Jest DOM matchers
import BankServicesList from "../components/BankServicesList/BankServicesList";
import { useGetBankServiceQuery } from "@/lib/redux/slices/bankService";

// Mock the hook
jest.mock("../lib/redux/slices/bankService", () => ({
  useGetBankServiceQuery: jest.fn(),
}));

// Type for the mock implementation
type MockedUseGetBankServiceQuery = jest.MockedFunction<
  typeof useGetBankServiceQuery
>;

describe("BankServicesList Component", () => {
  const mockedUseGetBankServiceQuery =
    useGetBankServiceQuery as MockedUseGetBankServiceQuery;

  beforeEach(() => {
    mockedUseGetBankServiceQuery.mockReset();
  });

  it("should display data correctly when fetch is successful", async () => {
    const mockData = {
      data: {
        data: {
          content: [
            {
              name: "Service 1",
              details: "Details about Service 1",
              numberOfUsers: 100,
              status: "Active",
              type: "Type A",
            },
          ],
        },
      },
      error: null,
      isLoading: false,
    };

    mockedUseGetBankServiceQuery.mockImplementation(() => mockData as any);

    render(<BankServicesList />);

    // Check for the content
    expect(await screen.findByText("Service 1")).toBeInTheDocument();
    expect(screen.getByText("Details about Service 1")).toBeInTheDocument();
  });
});

// it("should display loading state initially", () => {
//   mockedUseGetBankServiceQuery.mockImplementation(
//     () =>
//       ({
//         data: null,
//         error: null,
//         isLoading: true,
//       } as any)
//   ); // Use `as any` to bypass type errors

//   render(<BankServicesList />);
//   expect(screen.getByText("Bank Services List")).toBeInTheDocument();
//   // Adjust based on your actual loading indicator
//   expect(screen.getByText(/loading/i)).toBeInTheDocument();
// });

// it("should display error message when there is an error", async () => {
//   mockedUseGetBankServiceQuery.mockImplementation(
//     () =>
//       ({
//         data: null,
//         error: "An error occurred",
//         isLoading: false,
//       } as any)
//   );

//   render(<BankServicesList />);
//   expect(await screen.findByText("An error occurred")).toBeInTheDocument();
// });
