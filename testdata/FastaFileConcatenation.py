#paths to gene alignments 
geneAlignments = [r"path/to/genealignment",r"path/to/genealignment",r"path/to/genealignment"]

#path to new file
full_file = r"path/to/fullConcatenatedGenome.fasta"

#install biopython if you haven't already
from Bio import SeqIO

with open(full_file, 'w') as whole:
	for alignment in geneAlignments:
		with open(alignment) as part:
			records = SeqIO.parse(part, 'fasta')
			for record in records:
				SeqIO.write(record, whole, 'fasta')
