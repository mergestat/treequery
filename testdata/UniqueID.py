#START USER CONFIGURATION
#Store all sequence Ids that repeat in the sameIds array
#sameIds = ["id1","id2","id3"]

#path to original fasta file
#original_file = r"path/to/original/file.fasta"

#path to new file
#corrected_file = r"path/to/new/file.fasta"
#END USER CONFIGURATION

sameIds = ["South","Czech","Hong","New","Saudi","Sri",]

#path to original fasta file
original_file = r"COVIDGeneAlignments/ORF10_aligned.fasta"

#path to new file
corrected_file = r"correctedCOVIDGeneAlignments/corrected_ORF10_aligned.fasta"

#install biopython if you haven't already
from Bio import SeqIO

with open(original_file) as original, open(corrected_file, 'w') as corrected:
	records = SeqIO.parse(original_file, 'fasta')
	counter = 1
	for record in records:
	    for sameId in sameIds:
	        if record.id == "hCoV-19/" + sameId:
	        	#edit line below to alter id structure or counter
	            record.id = record.id + "/" + str(counter)
	            counter = counter + 1
	    print(record.id)
	    SeqIO.write(record, corrected, 'fasta')