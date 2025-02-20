"use client";

import { userStore } from "@/store/store";

export default () => {
  const { user } = userStore();
  return (
    <div className="w-full max-w-xl border shadow-lg rounded-xl mb-5 min-h-32 flex flex-col p-6 space-y-4 bg-white dark:bg-gray-900">
      <div className="flex flex-col items-center text-center space-y-2">
        <p className="text-2xl font-bold capitalize text-gray-900 dark:text-white">
          {user?.fullname}
        </p>
        <div className="flex flex-wrap justify-center items-center gap-2">
          <p className="text-sm text-gray-600 dark:text-gray-300">
            {user?.email}
          </p>
          <span className="px-3 py-1 text-sm font-medium border rounded-md bg-gray-100 dark:bg-gray-800 dark:text-gray-200">
            {user?.role}
          </span>
        </div>
      </div>
      <p className="text-center text-sm font-medium text-gray-500 dark:text-gray-400">
        BCA A Section IV Sem
      </p>
    </div>
  );
};
