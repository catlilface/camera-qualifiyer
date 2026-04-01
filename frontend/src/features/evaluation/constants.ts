export const MAX_FILE_SIZE = 50 * 1024 * 1024

export const methods = [
  {
    id: 'rr',
    title: 'RR',
    subtitle: 'Reference-based',
    description:
      'Метод с эталонным изображением для последующего сравнения.',
  },
  {
    id: 'nr',
    title: 'NR',
    subtitle: 'No-reference',
    description:
      'Метод без эталона, когда оценка строится только по загруженной картинке.',
  },
] as const

export const notes = [
  'Сейчас доступны только методы RR и NR.',
  'Выбор мониторов в этой версии не используется.',
  'Максимальный размер изображения: 50 МБ.',
] as const
