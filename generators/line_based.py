import argparse
import io
import subprocess

parser = argparse.ArgumentParser()
parser.add_argument('--structs', nargs='+')
parser.add_argument('--package')
parser.add_argument('--outfile')
args = parser.parse_args()

output = '''package %(package)s
''' % {'package': args.package}

for struct in args.structs:
    output += '''
    func (r *%(struct)s) Lines() []Line {
        return r.lines
    }

    func (r *%(struct)s) SetLines(lines []Line) {
        r.lines = lines
    }

    func (r *%(struct)s) AddLine(l Line) {
        if r.lines == nil {
            r.lines = make([]Line, 0, 10)
        }
        r.lines = append(r.lines, l)
    }

    ''' % {'struct': struct}

with io.open(args.outfile, 'w+', encoding='utf-8') as fp:
    fp.write(output)

subprocess.check_call(['gofmt', '-w', args.outfile])
