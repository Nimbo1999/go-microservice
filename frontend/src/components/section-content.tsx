import clsx from 'clsx'

type SectionContentProps = React.HTMLAttributes<HTMLDivElement>;

function SectionContent({ children, className, ...props }: SectionContentProps) {
  return (
    <div {...props} className={clsx([className, "section-content"])}>
      {props.title ? <h3 className="section-content__title">{props.title}</h3> : null}

      <div className="section-content__container">
        {children}
      </div>
    </div>
  )
}

export default SectionContent;
