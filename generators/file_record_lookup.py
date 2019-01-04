import argparse
import io
import subprocess

parser = argparse.ArgumentParser()
parser.add_argument('--recordtypes', nargs='+')
parser.add_argument('--package')
parser.add_argument('--outfile')
args = parser.parse_args()

output = '''package %(package)s
''' % {'package': args.package}

for typ in args.recordtypes:
    output += '''
    func (f *File) Lookup%(type)sByID(id string) (*%(type)sRecord, bool) {
        for _, r := range f.Records {
            t, tok := r.(*%(type)sRecord)
            i, iok := r.(Identifyable)
            if tok && iok && i.ID() == id {
                return t, true
            }
        }
        return nil, false
    }
    ''' % {'type': typ}

with io.open(args.outfile, 'w+', encoding='utf-8') as fp:
    fp.write(output)

subprocess.check_call(['gofmt', '-w', args.outfile])
