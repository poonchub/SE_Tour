import { createContext, ReactNode, useContext, useState } from 'react';

interface DateContextProps {
    dateSelectedFormat: any,
    setDateSelectedFormat: (param: string) => void,
    dateID: any,
    setDateID: (param: number) => void,
}

const DateContext = createContext<DateContextProps>({
    dateSelectedFormat: "",
    setDateSelectedFormat: () => {},
    dateID: 1,
    setDateID: () => {},
});

export const DateProvider = ({ children }: { children: ReactNode }) => {
    const [dateSelectedFormat, setDateSelectedFormat] = useState<string>("");
    const [dateID, setDateID] = useState<number>(1)
    return (
        <DateContext.Provider value={{ dateSelectedFormat, setDateSelectedFormat, dateID, setDateID }}>
            {children}
        </DateContext.Provider>
    );
};

export const useDateContext = (): DateContextProps => {
    const context = useContext(DateContext);
    if (!context) {
        throw new Error("useDateContext must be used within a DateProvider");
    }
    return context;
};