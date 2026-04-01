import { ImagePlus } from 'lucide-react'

type ImageUploadFieldProps = {
  error: string
  inputKey: number
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void
}

export function ImageUploadField({
  error,
  inputKey,
  onChange,
}: ImageUploadFieldProps) {
  return (
    <div className="space-y-3">
      <label className="text-sm font-medium text-slate-800">
        Изображение для анализа
      </label>
      <label className="flex cursor-pointer flex-col items-center justify-center rounded-[28px] border border-dashed border-sky-300 bg-sky-50/70 px-4 py-8 text-center transition hover:bg-sky-50 sm:px-6">
        <div className="flex h-14 w-14 items-center justify-center rounded-2xl bg-white text-sky-700 shadow-sm">
          <ImagePlus className="h-6 w-6" />
        </div>
        <p className="mt-4 text-base font-medium text-slate-900">
          Нажмите, чтобы выбрать файл
        </p>
        <p className="mt-2 max-w-md text-sm leading-6 text-slate-600">
          Поддерживаются изображения стандартных форматов. Размер файла не
          должен превышать 50 МБ.
        </p>
        <input
          key={inputKey}
          type="file"
          accept="image/*"
          className="sr-only"
          onChange={onChange}
        />
      </label>

      {error ? (
        <div className="rounded-2xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-700">
          {error}
        </div>
      ) : null}
    </div>
  )
}
