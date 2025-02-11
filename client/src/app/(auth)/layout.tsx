import React from "react";

const layout = ({ children }: { children: React.ReactNode }) => {
  return <div className="p-10">{children}</div>;
};

export default layout;
