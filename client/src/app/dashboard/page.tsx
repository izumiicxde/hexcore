"use client";

import { toast } from "@/hooks/use-toast";
import { subjectSummaryStore } from "@/store/store";
import { ISummaryAPIResponse } from "@/types/classes";
import { useEffect } from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  PieChart,
  Pie,
  Cell,
  Legend,
} from "recharts";

const COLORS = ["#10B981", "#EF4444"];

const CustomTooltip = ({ active, payload }: any) => {
  if (active && payload && payload.length) {
    return (
      <div className="p-2 rounded-lg shadow-md border border-gray-300 bg-white dark:bg-gray-800 dark:border-gray-700">
        <p className="text-sm font-semibold">
          {payload[0].name}: {payload[0].value}
        </p>
      </div>
    );
  }
  return null;
};

const Page = () => {
  const { summary, setSummary } = subjectSummaryStore();

  const getDetails = async () => {
    const url = `${process.env.NEXT_PUBLIC_API_ENDPOINT}/attendance/summary`;
    try {
      const response = await fetch(url, {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) {
        toast({ title: "Unexpected error while getting data." });
        return;
      }

      const data: ISummaryAPIResponse = await response.json();
      setSummary(data.summary);
    } catch (error: any) {
      toast({ title: "Unexpected error occurred" });
    }
  };

  useEffect(() => {
    if (!summary) getDetails();
  }, []);

  const barChartData = summary?.subjects
    ? Object.entries(summary.subjects).map(
        ([subject, data]: [string, any]) => ({
          subject,
          attended: data.attended_classes,
          remaining: data.remaining,
        })
      )
    : [];

  const pieChartData = summary
    ? [
        { name: "Attended", value: summary.total_attended },
        { name: "Missed", value: summary.total_missed },
      ]
    : [];

  return (
    <div className="lg:p-6 pt-10 w-full lg:max-w-4xl mx-auto lg:space-y-6">
      {summary ? (
        <>
          <div className="shadow-lg  rounded-xl p-6 flex flex-col items-center">
            <h1 className="text-3xl font-bold text-center">
              Attendance Summary
            </h1>
            <p className="text-lg font-semibold">
              Overall Percentage: {summary.overall_percentage}%
            </p>
            <p>
              Total Attended: {summary.total_attended} / {summary.total_classes}
            </p>
          </div>

          <div className="shadow-lg rounded-xl lg:p-6 p-1 pt-5">
            <h2 className="text-xl font-semibold mb-4 text-center">
              Attendance Overview
            </h2>
            <ResponsiveContainer width="100%" height={300}>
              <BarChart data={barChartData}>
                <XAxis dataKey="subject" className="text-sm" />
                <YAxis />
                <Tooltip
                  content={<CustomTooltip />}
                  cursor={{ fill: "transparent" }}
                />
                <Legend wrapperStyle={{ paddingTop: 10 }} />
                <Bar
                  dataKey="attended"
                  fill="#10B981"
                  name="Attended"
                  radius={[4, 4, 0, 0]}
                />
                <Bar
                  dataKey="remaining"
                  fill="#EF4444"
                  name="Remaining"
                  radius={[4, 4, 0, 0]}
                />
              </BarChart>
            </ResponsiveContainer>
          </div>

          <div className="shadow-lg rounded-xl p-6">
            <h2 className="text-xl font-semibold mb-4 text-center">
              Attendance Distribution
            </h2>
            <ResponsiveContainer width="100%" height={300}>
              <PieChart>
                <Pie
                  data={pieChartData}
                  dataKey="value"
                  nameKey="name"
                  cx="50%"
                  cy="50%"
                  outerRadius={100}
                  label
                >
                  {pieChartData.map((_, index) => (
                    <Cell
                      key={`cell-${index}`}
                      fill={COLORS[index % COLORS.length]}
                    />
                  ))}
                </Pie>
                <Legend wrapperStyle={{ paddingTop: 10 }} />
                <Tooltip content={<CustomTooltip />} />
              </PieChart>
            </ResponsiveContainer>
          </div>
        </>
      ) : (
        <p className="text-center">Loading attendance summary...</p>
      )}
    </div>
  );
};

export default Page;
