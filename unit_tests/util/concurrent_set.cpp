#include "util/concurrent_set.hpp"
#include "util/typedefs.hpp"

#include <boost/test/unit_test.hpp>
#include <set>

BOOST_AUTO_TEST_SUITE(concurrent_set_test)

using namespace osrm;
using namespace osrm::util;

BOOST_AUTO_TEST_CASE(concurrent_set_enable_func_test)
{
    ConcurrentSet<NodeID> ns1(false);
    BOOST_CHECK(!ns1.IsEnabled());

    ConcurrentSet<NodeID> ns2(true);
    BOOST_CHECK(ns2.IsEnabled());
}

BOOST_AUTO_TEST_CASE(node_set_disable_test)
{
    ConcurrentSet<NodeID> ns;
    BOOST_CHECK(!ns.IsEnabled());

    ns.Add(1);
    BOOST_CHECK(!ns.Check(1));
    BOOST_CHECK(!ns.Check(2));
}

BOOST_AUTO_TEST_CASE(node_set_enable_test)
{
    ConcurrentSet<NodeID> ns(true);
    BOOST_CHECK(ns.IsEnabled());

    ns.Add(1);
    BOOST_CHECK(ns.Check(1));

    BOOST_CHECK(!ns.Check(2));
    ns.Add(2);
    BOOST_CHECK(ns.Check(2));
}

BOOST_AUTO_TEST_CASE(node_set_percentage_test_empty_set)
{
    // Empty set, no TotalNum set
    ConcurrentSet<NodeID> ns1(true);
    std::string str1 = ns1.Statistic();
    BOOST_CHECK_EQUAL(str1, "There are 0 nodes be updated.  ");
}

BOOST_AUTO_TEST_CASE(node_set_percentage_test_none_empty_set)
{
    // None-empty set, no TotalNum set
    ConcurrentSet<NodeID> ns1(true);
    ns1.Add(1);
    ns1.Add(1);
    ns1.Add(2);
    std::string str1 = ns1.Statistic();
    BOOST_CHECK_EQUAL(str1, "There are 2 nodes be updated.  ");
}

BOOST_AUTO_TEST_CASE(node_set_percentage_test_totalnum_statistic)
{
    // None-empty set, with TotalNum set
    ConcurrentSet<NodeID> ns1(true);

    for (int i = 0; i < 23; ++i)
    {
        ns1.Add(i);
    }
    ns1.SetTotalNum(100);
    std::string str = ns1.Statistic();
    BOOST_CHECK_EQUAL(str, "There are 23 nodes be updated.  About 23% of 100 elements.\n");
}

BOOST_AUTO_TEST_CASE(node_set_percentage_test_rounding)
{
    // None-empty set, with TotalNum set
    ConcurrentSet<NodeID> ns1(true);

    for (int i = 0; i < 37; ++i)
    {
        ns1.Add(i);
    }
    ns1.SetTotalNum(84);
    std::string str = ns1.Statistic();
    BOOST_CHECK_EQUAL(str, "There are 37 nodes be updated.  About 44.05% of 84 elements.\n");
}

BOOST_AUTO_TEST_CASE(node_set_iteration)
{
    ConcurrentSet<NodeID> ns1(true);
    for (int i = 0; i < 10; ++i)
    {
        ns1.Add(i);
    }

    //int expect = 0;
    std::set<NodeID> s;
    for (auto iter : ns1)
    {
        s.insert(iter);
    }
    BOOST_CHECK_EQUAL(s.size(), 10);

}

BOOST_AUTO_TEST_SUITE_END()
