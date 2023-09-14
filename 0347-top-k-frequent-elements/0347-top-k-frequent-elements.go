func topKFrequent(nums[] int, k int)[] int {
    var counts = make(map[int] int)

    // cari angka unik beserta frekuensinya
    for i := 0; i < len(nums); i++ {
        counts[nums[i]] += 1;
    }

    var result = [] int {}
    for k > 0 {
        i := -1 // frekunsi pasti diatas 0
        var v int
        
        // cari nilai tertinggi
        for num, count := range counts {
            if i < count {
                i = count
                v = num
            }

        }
        
        // tampung dalam variabel,
        // kemudian hapus untuk menghindari redundansi
        result = append(result, v)
        delete(counts, v)
        k--;
    }
    return result
}