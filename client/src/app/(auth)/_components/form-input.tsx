import {
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

interface FormInputProps {
  name: string;
  label: string;
  type?: string;
  control: any;
}

export const FormInput = ({
  name,
  label,
  type = "text",
  control,
}: FormInputProps) => (
  <FormField
    control={control}
    name={name}
    render={({ field }) => (
      <FormItem>
        <FormLabel>{label}</FormLabel>
        <FormControl>
          <Input
            type={type}
            placeholder={`Enter your ${label.toLowerCase()}`}
            {...field}
            required
          />
        </FormControl>
        <FormMessage />
      </FormItem>
    )}
  />
);
