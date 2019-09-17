#ifndef OSRM_UTIL_CONCURRENT_SET_HPP
#define OSRM_UTIL_CONCURRENT_SET_HPP

#include "tbb/concurrent_unordered_set.h"

#include <set>
#include <string>
#include <sstream>
#include <iomanip>


namespace osrm
{
namespace util
{

template<typename T>
class ConcurrentSet
{
// https://www.threadingbuildingblocks.org/docs/help/reference/containers_overview/concurrent_unordered_set_cls.html
using SET = tbb::concurrent_unordered_set<T>;
using ITERATOR = typename tbb::concurrent_unordered_set<T>::iterator;
using CONST_ITERATOR = typename tbb::concurrent_unordered_set<T>::const_iterator;

public:
    ConcurrentSet(bool b = false)
    : enabled(b)
    {
        totalnum = -1;
    }

    void SetTotalNum(const int64_t num)
    {
        if (num <= 0)
        {
            return;
        }
        totalnum = num;
    }

    void Add(const T n)
    {
        if (enabled)
        {
           s.insert(n);
        }
     }

    void Clear()
    {
        if (enabled)
        {
            s.clear();
        }
        
    }

    bool Check(const T n) const
    {
        if (enabled)
        {
            return s.find(n) != s.end();
        }
        else
        {
            return false;
        }
    } 

    bool IsEnabled() const
    {
        return enabled;
    }

    std::string Statistic()
    {
        std::ostringstream ss;

        ss << "There are " << s.size() << " nodes be updated.  ";
        if (totalnum > 0)
        {
            float percentage = (float)(s.size() * 100) / totalnum;
            ss <<std::setprecision(4)<< "About " << percentage << "% of " << totalnum << " elements.\n";
        }

        return ss.str();
    }

    CONST_ITERATOR begin() const
    {
        return s.begin();
    }

    CONST_ITERATOR end() const
    {
        return s.end();
    }

private:
    SET         s;
    int64_t     totalnum;
    bool        enabled;
};

}

}

#endif
