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
    func (r *%(struct)s) ID() string {
        return r.id
    }

    func (r *%(struct)s) SetID(id string) {
        r.id = id
    }

    ''' % {'struct': struct}

with io.open(args.outfile, 'w+', encoding='utf-8') as fp:
    fp.write(output)

subprocess.check_call(['gofmt', '-w', args.outfile])
