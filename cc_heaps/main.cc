#include "max_priority_queue.hpp"
#include <string>
#include <iostream>
#include <cstdlib>

using namespace std;
using namespace myheaps;

struct ExtraData
{
    std::string description;
    int magic_number;

    ExtraData(const std::string &s, int m) : description(s), magic_number(m) {}
};

ostream &operator<<(ostream &out, const ExtraData &a)
{
    out << "{" << a.magic_number << "," << a.description << "}";
    return out;
}

void RandomCases()
{
    cout << endl
         << " *****************  Enter " << __func__ << endl;

    srand(time(NULL));

    MaxPriorityQueue<ExtraData> pq;

    //generate nodes
    cout << "build priority queue" << endl;
    for (int i = 0; i < 10; ++i)
    {
        assert(pq.IsMaxHeap());
        int n = rand() % 100;
        HeapNode<ExtraData> node(n, ExtraData(to_string(n * 100), n * 10));
        pq.Insert(node);

        cout << pq << endl;
    }

    cout << "extract from priority queue" << endl;
    while (!pq.empty())
    {
        assert(pq.IsMaxHeap());
        cout << "Max: " << pq.Maximum() << endl;

        HeapNode<ExtraData> node = pq.ExtractMax();
        cout << "ExtractMax: " << node << ", remain: " << pq << endl;
    }

    cout << " *****************  Exit " << __func__ << endl;
}

int main()
{

    MaxPriorityQueue<ExtraData> pq;
    cout << pq << endl; // dump heap to see change

    HeapNode<ExtraData> node1(10, ExtraData("n10", 10));
    pq.Insert(node1);
    cout << pq << endl;

    HeapNode<ExtraData> node2(9, ExtraData("n9", 9));
    int handle2 = pq.Insert(node2);
    cout << "handle2: " << handle2 << endl;
    cout << pq << endl;

    HeapNode<ExtraData> node3(15, ExtraData("n15", 15));
    int handle3 = pq.Insert(node3);
    cout << "handle3: " << handle3 << endl;
    cout << pq << endl;

    int new_handle2 = pq.IncreaseKey(handle2, 20);
    cout << "new_handle2: " << new_handle2 << endl;
    cout << pq << endl;

    int new_handle3 = pq.IncreaseKey(handle3, 21);
    cout << "new_handle3: " << new_handle3 << endl;
    cout << pq << endl;

    RandomCases();

    return 0;
}