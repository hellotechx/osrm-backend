#ifndef OSRM_CELLS_UPDATED_RECORD_HPP
#define OSRM_CELLS_UPDATED_RECORD_HPP

#include "partitioner/multi_level_partition.hpp"
#include "util/concurrent_set.hpp"


#include <vector>
#include <unordered_set>

namespace osrm
{
namespace customizer
{
namespace detail
{
class CellUpdateRecordImpl
{
public:
    CellUpdateRecordImpl(const partitioner::MultiLevelPartition& mlp)
    : partition(mlp),
      init(false)
    {
        std::vector<std::unordered_set<CellID>> tmp(partition.GetNumberOfLevels() - 1, std::unordered_set<CellID>());
        s = std::move(tmp);
    }

    void Collect(const util::ConcurrentSet<NodeID> &node_updated)
    {
        init = true;
        for (auto iter : node_updated)
        {
            for (std::size_t level = 1; level < partition.GetNumberOfLevels(); ++level)
            {
                s[level-1].insert(partition.GetCell(level, iter));
            }
        }
    }

    bool Check(LevelID l, CellID c) const
    {
        // always return true in default mode
        if (!init)
        {
            return true;
        }

        if (l < 1) 
        {
            printf("Incorrect level be passed.\n");
            return false;
        }
        return (s[l-1].find(c) != s[l-1].end());
    }

    void Clear()
    {
        for (int i = 0; i < partition.GetNumberOfLevels() - 1; ++i)
        {
            s[i].clear();
        }
        init = false;
    }

    std::string Statistic() const
    {
        return "\n";
    }

private:
    std::vector<std::unordered_set<CellID>> s;
    const partitioner::MultiLevelPartition &partition;
    bool init;
};
}

using CellUpdateRecord = detail::CellUpdateRecordImpl;

}
}

#endif
