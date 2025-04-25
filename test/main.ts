// compare-md5.ts
import { readFileSync } from 'fs';
// import overseas from "./wuwa-overseas-md5.json";
// import cnBilibili from "./wwSources-bilibili.json";

interface FileHash {
    path: string;
    hash: string;
}

function loadJsonFile(filename: string): FileHash[] {
    const content = readFileSync(filename, 'utf-8');
    return JSON.parse(content) as FileHash[];
}

function compareHashFiles(file1: string, file2: string) {
    const files1 = loadJsonFile(file1);
    const files2 = loadJsonFile(file2);

    // 创建路径到哈希的映射
    const map1 = new Map(files1.map(f => [f.path, f.hash]));
    const map2 = new Map(files2.map(f => [f.path, f.hash]));

    // 找出所有唯一的路径
    const allPaths = new Set([...map1.keys(), ...map2.keys()]);

    // 分类结果
    const identical: FileHash[] = [];
    const different: { path: string; hash1: string | null; hash2: string | null }[] = [];
    const onlyInFirst: FileHash[] = [];
    const onlyInSecond: FileHash[] = [];

    for (const path of allPaths) {
        const hash1 = map1.get(path);
        const hash2 = map2.get(path);

        if (hash1 && hash2) {
            if (hash1 === hash2) {
                identical.push({ path, hash: hash1 });
            } else {
                different.push({ path, hash1, hash2 });
            }
        } else if (hash1) {
            onlyInFirst.push({ path, hash: hash1 });
        } else {
            onlyInSecond.push({ path, hash: hash2! });
        }
    }

    return {
        identical,
        different,
        onlyInFirst,
        onlyInSecond,
        summary: {
            totalFiles: allPaths.size,
            identicalCount: identical.length,
            differentCount: different.length,
            onlyInFirstCount: onlyInFirst.length,
            onlyInSecondCount: onlyInSecond.length,
        }
    };
}


let file1 = "./wuwa-cn-bili-md5.json";
let file2 = "./wuwa-cn-md5.json";
file2 = "./wuwa-overseas-md5.json";

const result = compareHashFiles(file1, file2);

console.log('Comparison Results:');
console.log('-------------------');
console.log(`Total files: ${result.summary.totalFiles}`);
console.log(`Identical files: ${result.summary.identicalCount}`);
console.log(`Different files: ${result.summary.differentCount}`);
console.log(`Files only in ${file1}: ${result.summary.onlyInFirstCount}`);
console.log(`Files only in ${file2}: ${result.summary.onlyInSecondCount}`);

// 保存详细结果到文件
Bun.write('identical.json', JSON.stringify(result.identical, null, 2));
Bun.write('different.json', JSON.stringify(result.different, null, 2));
Bun.write(`only_in_${file1.split('/').pop()}.json`, JSON.stringify(result.onlyInFirst, null, 2));
Bun.write(`only_in_${file2.split('/').pop()}.json`, JSON.stringify(result.onlyInSecond, null, 2));

console.log('\nDetailed results saved to:');
console.log('- identical.json (path and hash match in both files)');
console.log('- different.json (path exists in both but hash differs)');
console.log(`- only_in_${file1.split('/').pop()}.json (files only in first file)`);
console.log(`- only_in_${file2.split('/').pop()}.json (files only in second file)`);
