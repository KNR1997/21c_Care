import cn from 'classnames';

interface useOpenAiProps {
  onClick: any;
  title: string;
  className?: string;
  isLoading?: boolean;
  disabled?: boolean;
}

export default function OpenAIButton({
  className,
  onClick,
  title,
  isLoading = false,
  disabled,
  ...rest
}: useOpenAiProps) {
  return (
    <button
      type="button"
      onClick={isLoading ? undefined : onClick} // Disable click when loading
      className={cn(
        'absolute right-0 -top-1 z-10 cursor-pointer text-sm font-medium text-accent hover:text-accent-hover',
        {
          'opacity-50 cursor-not-allowed': isLoading, // Visual feedback for loading state
          'hover:text-accent-hover': !isLoading,
        },
        className,
      )}
      disabled={disabled}
      {...rest}
    >
      {isLoading ? (
        <div className="flex items-center gap-1">
          <svg
            className="animate-spin h-4 w-4"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              className="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              strokeWidth="4"
            />
            <path
              className="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            />
          </svg>
          <span>{title}</span>
        </div>
      ) : (
        title
      )}
    </button>
  );
}
