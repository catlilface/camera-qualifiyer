import { MAX_FILE_SIZE } from './constants'

export function formatFileSize(size: number) {
  if (size < 1024 * 1024) {
    return `${(size / 1024).toFixed(1)} КБ`
  }

  return `${(size / (1024 * 1024)).toFixed(2)} МБ`
}

export function validateImageFile(file: File) {
  if (!file.type.startsWith('image/')) {
    return 'Можно загружать только изображения.'
  }

  if (file.size > MAX_FILE_SIZE) {
    return 'Файл превышает ограничение 50 МБ. Выберите изображение меньше.'
  }

  return ''
}

export function revokePreviewUrl(url: string | null) {
  if (url) {
    URL.revokeObjectURL(url)
  }
}
