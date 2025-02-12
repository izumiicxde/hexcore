import React from "react";

const layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="p-10 flex justify-center items-center h-screen overflow-hidden">
      {children}
    </div>
  );
};

export default layout;
